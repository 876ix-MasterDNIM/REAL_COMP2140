$(document).ready(function(){
   $.validator.addMethod('phoneNumberCheck', function(value, element){
      return /^(\(876\)[\s]?|^1?876[\s-]?)?\d{3}[\s-]?\d{4}$/.test(value);
  }, 'Not a valid Jamaican phone number');

  $('#purchaseform').validate({
      rules: {
         first: 'required',
         last: 'required',
         cname: 'required',
         cnumber:{
            required: true,
            phoneNumberCheck: true
         },
         startdate: 'required',
         enddate: 'required'
         //size:
      }, // end rules
      messages: {
         first: 'First name is required',
         last: 'Last name is required',
         cname: 'Company name is required',
         cnumber: {
            required: 'Company number is required',
         },
         startdate: 'Please choose a date',
         enddate: 'please choose a date',
      }, // end messages
      errorPlacement: function(error, element) {
          if ( element.is(':radio') || element.is(':checkbox')) {
             if (!element.hasClass('input-field')){
               error.appendTo( element.parent());
            }
          } else {
            error.insertAfter(element);
          }
      } // end errorPlacement
   });
});
