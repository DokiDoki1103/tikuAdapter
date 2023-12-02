rm -rf cmd/adapter-service/dist
cd web || exit
npm install && npm run build
mv -f dist ../cmd/adapter-service