<html>
<head>
<tile></tile>
</head>
<body>

<from action="/login" method="post">
    <input type="checkbox" name="interest" value="football">足球
    <input type="checkbox" name="interest" value="basketball">篮球
    <input type="checkbox" name="interest" value="tennis">网球
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登陆">

</from>

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
