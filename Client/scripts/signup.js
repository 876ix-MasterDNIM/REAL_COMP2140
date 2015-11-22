$(document).ready(function() {
   var validJcanPhoneNum = /^(\(876\)[\s]?|^1?876[\s-]?)?\d{3}[\s-]?\d{4}$/; // idk regex? lel.

   $('#cctype').on('change', function(){
      $(this).focusout();
   });

   // $.validator.addMethod('valueNotEquals', function(value, element, param) {
   //   return (value == param);
   // }, 'You haven\'t selected a credit card');

   $.validator.addMethod('phoneNumberCheck', function(value, element) {
      return validJcanPhoneNum.test(value);
   }, 'Not a valid Jamaican phone number');

   $('#signupform').validate({
      rules: {
         first: 'required',
         last: 'required',
         email: {
            required: true,
            email: true
         },
         username: 'required',
         password: {
            required: true,
            rangelength: [6, 12]
         },
         cpassword: {
            equalTo: '#password'
         },
         cc: {
            required: true,
            creditcard: true
         },
         cvv: {
            required: true,
            minlength: 3,
            number: true
         },
         ctype: {
            required: true,
            // valueNotEquals: 'Please choose a credit card'
         },
         telephone: {
            required: true,
            phoneNumberCheck: true
         },
         company: 'required'
      }, // end rules
      messages: {
         first: 'Please enter your firstname',
         last: 'Please enter your lastname',
         email: 'Please enter a valid email address',
         username: 'Please enter the user name you\'d like to use',
         password: {
            required: 'Please provide a password',
            rangelength: 'Your password must 6-12 characters long'
         },
         cpassword: {
            equalTo: 'Passwords must match'
         },
         cc: {
            required: 'Please provide a credit card number',
            creditcard: 'Not a valid credit card number'
         },
         cvv: {
            required: 'Please enter your card\'s CVV',
            minlength: 'A valid CVV is at least 3 digits long',
            number: 'Value is not a number'
        },
         ctype: {
            required: 'Must select a credit card'
        },
         telephone: {
            required: 'Please enter your phone number',
         },
         company: 'Please enter your company\'s name'
      }
   }); //end validate
});
