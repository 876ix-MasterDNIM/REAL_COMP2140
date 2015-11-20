$('document').ready(function() {
   $('.login-form').validate({
      rules: {
         username: 'required',
         password: {
            required: true,
            rangelength: [6, 12]
         }
      },// end rules
      messages: {
         username: "Username is required",
         password: {
            required: 'password is required',
            rangelength: 'must be 6-12 characters long'
         }
      }
  });
});
