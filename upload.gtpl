<html>
<head>
       <title>Upload file</title>
</head>
<body>
<a href="/home">home</a>
<form enctype="multipart/form-data" action="upload" method="post">
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}"/>
  folder:<br>
  <input type="text" name="folder"><br>

    <input type="submit" value="upload" />
</form>
</body>
</html>
