import { Component, OnInit } from '@angular/core'
import {FormBuilder, FormGroup, Validators} from '@angular/forms'
import { en_US, NzI18nService} from 'ng-zorro-antd/i18n'

@Component({
  selector: 'dod-task-form',
  templateUrl: './task-form.component.html',
  styleUrls: ['./task-form.component.scss']
})
export class TaskFormComponent implements OnInit {

  validateForm!: FormGroup

  submitForm(): void {
    for (const i in this.validateForm.controls) {
      this.validateForm.controls[i].markAsDirty()
      this.validateForm.controls[i].updateValueAndValidity()
    }
  }

  constructor(private fb: FormBuilder,
              private i18n: NzI18nService) {}

  ngOnInit(): void {
    this.i18n.setLocale(en_US)
    this.validateForm = this.fb.group({
      userName: [null, [Validators.required]],
      password: [null, [Validators.required]],
      remember: [true]
    })
  }

}
