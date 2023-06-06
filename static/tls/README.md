## tls

    * 私钥key.pem
    * 证书cert.pem

### 用内置工具

    ```
    ./releases/tls
    ```

### ~~用openssl~~

    ```
    openssl genrsa -out static/tls/key.pem 2048
    openssl req -new -x509 -key static/tls/key.pem -out static/tls/cert.pem -days 3650
    ```