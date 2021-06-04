import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'sortArrayByProp'
})
export class SortArrayByPropPipe implements PipeTransform {
  transform(records: Array<any>, args?: any): any {
    return records.sort(function(a, b){
      if(a[args.property] < b[args.property]){
        return -1 * args.direction;
      }
      else if( a[args.property] > b[args.property]){
        return 1 * args.direction;
      }
      else{
        return 0;
      }
    });
  };
}
