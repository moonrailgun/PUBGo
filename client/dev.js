const { createProxyMiddleware } = require('http-proxy-middleware');
const Bundler = require('parcel-bundler');
const express = require('express');

const bundler = new Bundler('index.html');
const app = express();
const port = Number(process.env.PORT || 1234);

app.use(
  '/api',
  createProxyMiddleware({
    target: 'http://localhost:9091',
    changeOrigin: true,
    pathRewrite: { '^/api': '' },
  })
);

console.log(`正在监听端口:${port}`);
app.use(bundler.middleware());

app.listen(port);
