import {ToastrService} from 'ngx-toastr';

export function pushErrorNotification(err, toastr: ToastrService): void {
  if (toastr != null) {
    toastr.error(err.toString(), 'Error Message', {
      timeOut: 3000,
    });
  } else {
    console.error(
      'An Error occurred and could not be displayed via toast! Details => ' +
      JSON.stringify(err)
    );
  }
}

export function pushInfoNotification(msg, toastr: ToastrService): void {
  if (toastr != null) {
    toastr.info(msg, 'Information', {
      timeOut: 3000,
    });
  } else {
    console.log(
      'An informational event/msg was received and could not be displayed via toast! Details => ' +
      JSON.stringify(msg)
    );
  }
}
