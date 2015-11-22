$(document).ready(function() {
   var validJcanPhoneNum = /^(\(876\)[\s]?|^1?876[\s-]?)?\d{3}[\s-]?\d{4}$/;
   var validDate = /^\d{1,2}\s[a-zA-Z]{3,9},\s\d{4}$/; // regex for validDate.

   var file = $('#file');
   var isUploadBtnHidden = false;

   $('#yes').click(function() {
      if (!isUploadBtnHidden) {
      file.prop('disabled', true);
      $('#filespan').hide(300);
   }
      isUploadBtnHidden = true;
   });

   $('#no').click(function() {
      if (isUploadBtnHidden) {
         file.prop('disabled', false);
         $('#filespan').show(250);
      }
      isUploadBtnHidden = false;
   });

   $('.datepicker').pickadate({
      // was tryna make it so you can only set a date 3 months into the future , aint working tho.
      //min: new Date(today.getFullYear(), today.getMonth(), today.getDay()),
      //max: new Date(today.getFullYear(), today.getMonth() + 1, today.getDay()),
      selectMonths: true, // Creates a dropdown to control month
      selectYears: 15 // Creates a dropdown of 15 years to control year
   });

   // start and enddate
   // $('#startdate').on('changeDate', function() {
   //   console.log("Date changed");
   // });
   // $('#enddate').on('changeDate', function() {
   //   console.log('enddate changed');
   // });


   $.validator.addMethod('valueNotEquals', function(value, element, param) {
      return (value == param);
   }, 'You haven\'t selected a credit card');

   $.validator.addMethod('phoneNumberCheck', function(value, element) {
      return validJcanPhoneNum.test(value);
  }, 'Not a valid Jamaican phone number');

  $.validator.addMethod('dateCheck', function(value, element) {
     console.log('value is ' + value);
     return validDate.test(value);
  }, 'Please choose a date');

   $('#purchaseform').validate({
      rules: {
         first: 'required',
         last: 'required',
         cname: 'required',
         cnumber:{
            required: true,
            phoneNumberCheck: true
         },
         startdate: {
            //required: true,
            dateCheck: true
         },
         enddate: {
            //required: true,
            dateCheck: true
         },
         size: {
            valueNotEquals: 'Size'
         }
      }, // end rules
      messages: {
         first: 'First name is required',
         last: 'Last name is required',
         cname: 'Company name is required',
         cnumber: {
            required: 'Company number is required',
         },
         // startdate: 'Please choose a date',
         // enddate: 'please choose a date',
         size: {
            required: true,
            valueNotEquals: 'Please choose a size'
         }
      }, // end messages
      errorPlacement: function(error, element) {
          if ( element.is(':radio') || element.is(':checkbox')) {
             if (!element.hasClass('input-field')){
               error.appendTo( element.parent());
            }
          } else {
            error.insertAfter(element);
          }
      } // end errorPlacement Still not sure whether this method relle tbh.
   });
});
