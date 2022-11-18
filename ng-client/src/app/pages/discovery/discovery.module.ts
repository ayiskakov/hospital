import { NgModule } from '@angular/core';

import { DiscoveryRoutingModule } from './discovery-routing.module';

import { NzTableModule } from "ng-zorro-antd/table";
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { ReactiveFormsModule } from "@angular/forms";
import { NzFormModule } from "ng-zorro-antd/form";
import { NzInputModule } from "ng-zorro-antd/input";
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker';
import { CommonModule } from "@angular/common";
import { DiscoveryComponent } from './discovery.component';


@NgModule({
  imports: [
    DiscoveryRoutingModule,
    NzTableModule,
    NzButtonModule,
    NzModalModule,
    ReactiveFormsModule,
    NzFormModule,
    NzInputModule,
    NzInputNumberModule,
    NzIconModule,
    NzDatePickerModule,
    NzSelectModule,
    CommonModule,
  ],
  declarations: [DiscoveryComponent],
  exports: [DiscoveryComponent],
})
export class DiscoveryModule { }
