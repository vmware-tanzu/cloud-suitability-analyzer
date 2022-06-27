import {Component, EventEmitter, Input, OnChanges, Output, SimpleChanges} from '@angular/core';
import {Findings} from '../../../../model/findings';

@Component({
  selector: 'app-finding-details',
  templateUrl: './finding-details.component.html',
  styleUrls: ['./finding-details.component.css']
})
export class FindingDetailsComponent implements OnChanges {

  constructor() {
  }

  ngOnChanges(changes: SimpleChanges): void {
  }

  @Input()
  selectedFinding: Findings;

  @Input()
  isOpen: boolean;

  @Output()
  isOpenChange = new EventEmitter();

  public onModalClose(opened: boolean): void {
    this.isOpenChange.emit(opened);
  }
}
