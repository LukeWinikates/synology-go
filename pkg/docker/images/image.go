package images

//http://192.168.1.33:5000/webapi/entry.cgi/SYNO.Docker.Image
//api=SYNO.Docker.Image&method=get&version=1&image=%22hello-world%22&tag=%22latest%22

// api=SYNO.Docker.Image&method=upgrade_start&version=1&repository=%22hello-world%22
// response:
// {"data":{"task_id":"@administrators/SYNO_DOCKER_IMAGE_UPGRADE17162526376711FFDA"},"success":true}

// api=SYNO.Docker.Image&method=upgrade_status&version=1&task_id=%22%40administrators%2FSYNO_DOCKER_IMAGE_UPGRADE17162526376711FFDA%22
// response:
// {"data":{"current":0,"finished":false,"image":"hello-world:latest","state":"PULLING_IMAGE","total":0},"success":true}

// api=SYNO.Docker.Image&method=get&version=1&image=%22lukewinikates%2Forbi-exporter%22&tag=%22latest%22
// {"data":{"author":"","cmd":[],"digest":"sha256:26973ca9d9a7a093dd2ff939337caa0785cc97db5f0d9c08c96756ecc454dd7d","docker_version":"","entrypoint":["/orbi-exporter-linux-amd64"],"env":[{"key":"PATH","value":"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},{"key":"SSL_CERT_FILE","value":"/etc/ssl/certs/ca-certificates.crt"}],"id":"sha256:374367a73678b2097724791be28b8175e1d7b2935469dfd4c2773d7c6bd710e9","image":"lukewinikates/orbi-exporter","ports":[{"port":"6724","protocol":"tcp"}],"size":13560874,"tag":"latest","virtual_size":13560874,"volumes":[]},"success":true}
