import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'searchFilter'
})
export class SearchFilterPipe implements PipeTransform {

  transform(value: any, param: any, args?: any): any {
    /* istanbul ignore else */
    if (!value) {
      return null;
    }
    /* istanbul ignore else */
    if (!args) {
      return value;
    }

    args = args.toLowerCase();
    return value.filter(function (data) {
      return JSON.stringify(data).toLowerCase().includes(args);
    });
  }

}
