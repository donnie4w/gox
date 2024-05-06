# gox

- The design philosophy of Go language promotes the crafting of clean and succinct code. A prominent characteristic of Go is its simplicity and efficiency, empowering developers to create robust and high-performance applications with notably less code compared to some other languages.

- gox is a static resource server, with 30 lines of code, you can achieve a support https,http and other services of static resource services.

Since http(s) services are often started in the system for js or some static resource testing. nginx is still relatively troublesome, so use go to write a static resource service gox, with a few lines of code can achieve a good performance, and stable service. As long as you run gox, you can use http(s) to access all files in the same directory. No configuration is required. It's more convenient.

- ./gox
    - start the http service by default, port 8080, and print the access path

- ./gox -port=8000 -https -nolog
    - -port    listening port
    - -https   starting the https Service
    - -nolog   the access path log is not printed

------------

[Download the gox execution file](https://tlnet.top/download "Download the gox execution file")

![](https://tlnet.top/statics/tlnet/115643.gif)
