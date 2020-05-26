import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

/** Modules */
import {UserModule} from './modules/user/user.module';

import { AppComponent } from './app.component';
import {SharedModule} from './modules/shared/shared.module';
import {HttpClientModule} from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    UserModule,
    SharedModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
