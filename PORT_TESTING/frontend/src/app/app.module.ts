import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';

import {HomeComponent} from './PAGES/home/home.component'
import {Page1Component} from './PAGES/page1/page1.component'
import {Page2Component} from './PAGES/page2/page2.component'
import {Page3Component} from './PAGES/page3/page3.component'

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

//allows communication with a remote server over HTTP (w/ GO)
import { HttpClientModule } from '@angular/common/http'; 

//add your declarations and imports here
@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    Page1Component,
    Page2Component,
    Page3Component,
    Page1Component,
    Page2Component,
    Page3Component,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    RouterModule,

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
