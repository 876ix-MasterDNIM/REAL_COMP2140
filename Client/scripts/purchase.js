$(document).ready(function() {
   var validJcanPhoneNum = /^(\(876\)[\s]?|^1?876[\s-]?)?\d{3}[\s-]?\d{4}$/;
   var validDate = /^\d{1,2}\s[a-zA-Z]{3,9},\s\d{4}$/; // regex for validDate.

   var file = $('#file');
   var isUploadBtnHidden = false;
   var sizeChangedOnce = false;

   $('.datepicker').on('change', function() {
      $(this).focusout();
      $(this).focusin();
   });

   $('#yes').on('click', function() {
      if (!isUploadBtnHidden) {
      file.prop('disabled', true);
      $('#upload').hide(300);
   }
      isUploadBtnHidden = true;
   });

   $('#no').on('click', function() {
      if (isUploadBtnHidden) {
         file.prop('disabled', false);
         $('#upload').fadeIn(250);
      }
      isUploadBtnHidden = false;
   });

   $('#size').on('change', function(){
      $(this).focusout();
   });

   $('.datepicker').pickadate({
      // was tryna make it so you can only set a date 3 months into the future , aint working tho.
      // min: new Date(today.getFullYear(), today.getMonth(), today.getDay()),
      // max: new Date(today.getFullYear(), today.getMonth() + 1, today.getDay()),
      selectMonths: true, // Creates a dropdown to control month
      selectYears: 15 // Creates a dropdown of 15 years to control year
   });

   $.validator.addMethod('valueNotEquals', function(value, element, param) {
      return (value != param);
   }, 'Please choose a size');

   $.validator.addMethod('phoneNumberCheck', function(value, element) {
      return validJcanPhoneNum.test(value);
  }, 'Not a valid Jamaican phone number');

  $.validator.addMethod('dateCheck', function(value, element) {
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
            // required: true,
            dateCheck: true
         },
         enddate: {
            // required: true,
            dateCheck: true
         },
         day: {
            required: true
         },
         size: {
            valueNotEquals: true,
            required: true
         },
         color: {
            required: true
         }
      }, // end rules
      messages: {
         first: 'First name is required',
         last: 'Last name is required',
         cname: 'Company name is required',
         cnumber: {
            required: 'Company number is required',
         },
         size: {
            required: "Please chose a size"
         }
      }, // end messages
      errorPlacement: function(error, element) {
          if ( element.is(':radio') || element.is(':checkbox')) {
             // console.log(element.parent().attr('class'));
             error.appendTo(element.parent());
          } else {
            error.insertAfter(element);
          }
      } // end errorPlacement
   });
});
