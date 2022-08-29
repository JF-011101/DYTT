

# sudo tee -a /etc/hosts <<EOF
# 127.0.0.1 dytt.api.com
# 127.0.0.1 dytt.comment.com
# 127.0.0.1 dytt.favorit.com
# 127.0.0.1 dytt.feed.com
# 127.0.0.1 dytt.publish.com
# 127.0.0.1 dytt.relation.com
# 127.0.0.1 dytt.user.com
# EOF


# 生成证书和私钥
cd ${DYTT_ROOT}/config/cert/api/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt api-csr.json | cfssljson -bare dytt-api



cd ${DYTT_ROOT}/config/cert/comment/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt comment-csr.json | cfssljson -bare dytt-comment

cd ${DYTT_ROOT}/config/cert/favorite/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt favorite-csr.json | cfssljson -bare dytt-favorite

cd ${DYTT_ROOT}/config/cert/feed/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt feed-csr.json | cfssljson -bare dytt-feed

cd ${DYTT_ROOT}/config/cert/publish/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt publish-csr.json | cfssljson -bare dytt-publish

cd ${DYTT_ROOT}/config/cert/relation/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt relation-csr.json | cfssljson -bare dytt-relation

cd ${DYTT_ROOT}/config/cert/user/
cfssl gencert -ca=${DYTT_ROOT}/config/cert/ca.pem \
  -ca-key=${DYTT_ROOT}/config/cert/ca-key.pem \
  -config=${DYTT_ROOT}/config/cert/ca-config.json \
  -profile=dytt user-csr.json | cfssljson -bare dytt-user