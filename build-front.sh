#!/bin/bash

cd app
npm ci --legacy-peer-deps
npm exec svelte-kit sync
npm run build
cp -rf ../build/ ./build