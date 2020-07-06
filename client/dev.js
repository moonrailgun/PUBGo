const { createProxyMiddleware } = require('http-proxy-middleware');
const Bundler = require('parcel-bundler');
const express = require('express');

let bundler = new Bundler('index.html');
let app = express();

app.use(
  '/api',
  createProxyMiddleware({
    target: 'http://localhost:9091',
    changeOrigin: true,
    pathRewrite: { '^/api': '' },
  })
);

app.use(bundler.middleware());

app.listen(1234);
