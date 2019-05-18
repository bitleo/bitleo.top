 $(function () {
   // $("#loading-test").click(function() {
   //      var $this = $(this);
   //      var loadingText = '<i class="fa fa-circle-o-notch fa-spin"></i> loading...';
   //      if ($(this).html() !== loadingText) {
   //        $this.data('original-text', $(this).html());
   //        $this.html(loadingText);
   //      }
   //       setTimeout(function() {
   //         $this.html($this.data('original-text'));
   //        }, 2000);
   //  });
     
    $("#convert-btn").click(function () {
        var $this = $(this);
        var loadingText = '<i class="fa fa-circle-o-notch fa-spin"></i> convert...';
        if ($(this).html() !== loadingText) {
          $this.data('original-text', $(this).html());
          $this.html(loadingText);
        }
        var file = $("#submitFile").get(0).files[0];
        if(file){
            var fileSize = file.size;
            if(fileSize>1*1024*1024){
                $this.html($this.data('original-text'));
                alert("can not upload file >1MB")
                return;
            }
            $("#convert-file-id").ajaxSubmit({
            success: function (data) {
               // alert(data.data);
                $this.html($this.data('original-text'));
                window.open("/pdfconverttoword/getfile?id="+data.data,"_blank")
            },
            error: function (error) { 
                $this.html($this.data('original-text'));
                alert("covert failed");
            },
            url: '/pdfconverttoword/convert', /*设置post提交到的页面*/
            type: "post", /*设置表单以post方法提交*/
            dataType: "json" /*设置返回值类型为文本*/
          });     
        }else{
            $this.html($this.data('original-text'));
            alert("upload file is null");
        }
    });
});