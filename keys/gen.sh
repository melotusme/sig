openssl genrsa -out ca.key 2048
openssl rsa -in ca.key -pubout -out openssl_pub.key

