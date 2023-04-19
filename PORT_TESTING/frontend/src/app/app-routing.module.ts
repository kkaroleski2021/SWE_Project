import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {HomeComponent} from './PAGES/home/home.component'
import {Page1Component} from './PAGES/page1/page1.component'
import {Page2Component} from './PAGES/page2/page2.component'
import {Page3Component} from './PAGES/page3/page3.component'

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'page1', component:Page1Component},
  {path: 'page2', component:Page2Component},
  {path: 'page3', component:Page3Component},

  {path:"**", redirectTo: ''}, // otherwise redirect to home
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
