import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { FormsModule, ReactiveFormsModule } from '@angular/forms'

import { IconsProviderModule } from '../icons-provider/icons-provider.module'

import { TaskFormComponent } from './components/task-form/task-form.component'

import {
  NzButtonModule,
  NzDatePickerModule,
  NzInputModule,
  NzFormModule,
} from 'ng-zorro-antd'

@NgModule({
  declarations: [TaskFormComponent],
  exports: [
    TaskFormComponent
  ],
  imports: [
    CommonModule,
    IconsProviderModule,
    FormsModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzDatePickerModule,
    NzInputModule,
    NzFormModule
  ]
})
export class SharedModule { }
