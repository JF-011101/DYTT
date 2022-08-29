mkdir -p $HOME/bin/
wget https://github.com/cloudflare/cfssl/releases/download/v1.6.1/cfssl_1.6.1_linux_amd64 -O $HOME/bin/cfssl
wget https://github.com/cloudflare/cfssl/releases/download/v1.6.1/cfssljson_1.6.1_linux_amd64 -O $HOME/bin/cfssljson
wget https://github.com/cloudflare/cfssl/releases/download/v1.6.1/cfssl-certinfo_1.6.1_linux_amd64 -O $HOME/bin/cfssl-certinfo
#wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -O $HOME/bin/cfssl
#wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -O $HOME/bin/cfssljson
#wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64 -O $HOME/bin/cfssl-certinfo
chmod +x $HOME/bin/{cfssl,cfssljson,cfssl-certinfo}