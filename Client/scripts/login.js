$(document).ready(function() {
   $('#login').validate({
      rules: {
         username: {
            required: true,
            // remote: can use this to call the server to validate but caah botha tbh.
         },
         password: {
            required: true,
            rangelength: [6, 12]
         }
      },// end rules
      messages: {
         username: "Username is required",
         password: {
            required: 'Password is required',
            rangelength: 'must be 6-12 characters long'
         }
      }
  });
});
