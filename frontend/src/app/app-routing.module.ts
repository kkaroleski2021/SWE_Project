import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {CComponent} from './login/login.component';
import {AComponent} from './home/home.component';
import { SignupComponent } from './signup/signup.component';


const routes: Routes = [
  {
    path: 'c-component',
    component: CComponent
  },
  {
    path: '',
    component: AComponent
  },
  {
    path: "signup",
    component: SignupComponent
  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
