## CA证书配置

为了确保安全，DYTT 系统各服务组件需要使用 x509 证书对通信进行加密和认证。所以，这里我们需要先创建 CA 证书。CA 根证书是所有组件共享的，只需要创建一个 CA 证书，后续创建的所有证书都由它签名。我们可以使用 CloudFlare 的 PKI 工具集 cfssl 来创建所有的证书。安装 cfssl 工具集。我们可以直接安装 cfssl 已经编译好的二进制文件，cfssl 工具集中包含很多工具，这里我们需要安装 cfssl、cfssljson、cfssl-certinfo，功能如下。
- cfssl：证书签发工具。
- cfssljson：将 cfssl 生成的证书（json 格式）变为文件承载式证书。

这两个工具的安装方法如下：
```bash
sh $DYTT_ROOT/scripts/cfsslinstall.sh
```

创建配置文件。CA 配置文件是用来配置根证书的使用场景 (profile) 和具体参数 (usage、过期时间、服务端认证、客户端认证、加密等)，可以在签名其它证书时用来指定特定场景：

```bash
$ cd $DYTT_ROOT/config/cert
$ tee ca-config.json << EOF
{
    "signing": {
      "default": {
        "expiry": "87600h"
      },
      "profiles": {
        "dytt": {
          "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
          ],
          "expiry": "876000h"
        }
      }
    }
  }
EOF
```
上面的 JSON 配置中，有一些字段解释如下。signing：表示该证书可用于签名其它证书（生成的 ca.pem 证书中 CA=TRUE）。server auth：表示 client 可以用该证书对 server 提供的证书进行验证。client auth：表示 server 可以用该证书对 client 提供的证书进行验证。expiry：876000h，证书有效期设置为 100 年。

创建证书签名请求文件。我们创建用来生成 CA 证书签名请求（CSR）的 JSON 配置文件：
```bash
$ cd $DYTT_ROOT/config/cert
$ tee ca-csr.json << EOF
{
    "CN": "dytt-ca",
    "key": {
      "algo": "rsa",
      "size": 2048
    },
    "names": [
      {
        "C": "CN",
        "ST": "GuangDong",
        "L": "ShenZhen",
        "O": "jf-011101",
        "OU": "dytt"
      }
    ],
    "ca": {
      "expiry": "876000h"
    }
  }
EOF
```
上面的 JSON 配置中，有一些字段解释如下。C：Country，国家。ST：State，省份。L：Locality (L) or City，城市。CN：Common Name，iam-apiserver 从证书中提取该字段作为请求的用户名 (User Name) ，浏览器使用该字段验证网站是否合法。O：Organization，iam-apiserver 从证书中提取该字段作为请求用户所属的组 (Group)。OU：Company division (or Organization Unit – OU)，部门 / 单位。

除此之外，还有两点需要我们**注意**。不同证书 csr 文件的 CN、C、ST、L、O、OU 组合必须不同，否则可能出现 PEER'S CERTIFICATE HAS AN INVALID SIGNATURE 错误。后续创建证书的 csr 文件时，CN、OU 都不相同（C、ST、L、O 相同），以达到区分的目的。

创建 CA 证书和私钥 首先，我们通过 cfssl gencert 命令来创建：
```bash
$ cd $DYTT_ROOT/config/cert
$ cfssl gencert -initca ca-csr.json | cfssljson -bare ca
```
上述命令会创建运行 CA 所必需的文件 ca-key.pem（私钥）和 ca.pem（证书），还会生成 ca.csr（证书签名请求），用于交叉签名或重新签名。创建完之后，我们可以通过 cfssl certinfo 命名查看 cert 和 csr 信息：
```bash
$ cfssl certinfo -cert ${IAM_CONFIG_DIR}/cert/ca.pem # 查看 cert(证书信息)
$ cfssl certinfo -csr ${IAM_CONFIG_DIR}/cert/ca.csr # 查看 CSR(证书签名请求)信息
```

配置 hosts：
```bash
sudo tee -a /etc/hosts <<EOF
127.0.0.1 dytt.api.com
127.0.0.1 dytt.comment.com
127.0.0.1 dytt.favorit.com
127.0.0.1 dytt.feed.com
127.0.0.1 dytt.publish.com
127.0.0.1 dytt.relation.com
127.0.0.1 dytt.user.com
EOF
```

### 各服务证书和私钥的创建


先创建各服务证书签名请求：


比如：
```bash
$ cd $DYTT_ROOT/config/cert/user
$ tee user-csr.json <<EOF
{
    "CN": "dytt-user",
    "key": {
      "algo": "rsa",
      "size": 2048
    },
    "names": [
        {
          "C": "CN",
          "ST": "GuangDong",
          "L": "ShenZhen",
          "O": "jf-011101",
          "OU": "dytt"
        }
    ],
    "hosts": [
      "127.0.0.1",
      "localhost",
      "dytt.user.com"
    ]
}
EOF
```

创建完后执行 ./scripts/certgen.sh 脚本生成各服务的证书和私钥：
```bash
$ sh $DYTT_ROOT/scripts/certgen.sh
```

## 参考
极客时间课程--《Go语言项目开发实战》-孔令飞