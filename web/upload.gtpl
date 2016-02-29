<html>
<head>
<tile></tile>
</head>
<body>

<form enctype="multipart/form-data" action="http://127.0.0.1:9090/upload" method="post">
  <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="upload" />
      </form>
<script>


alert("hello")
</script>

</body>
</html>
