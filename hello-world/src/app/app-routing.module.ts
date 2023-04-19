import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListingPageComponent } from './listing-page/listing-page.component';
import { ProfilesComponent } from './profiles/profiles.component';
import { SellerPageComponent } from './seller-page/seller-page.component';
import { SComponent } from './signup/signup.component';
import { BComponent} from './login/login.component';

const routes: Routes = [
  {
    path: '',
    component: ProfilesComponent   
  },
  {
    path: 'Login',
    component: BComponent
  },
  {
    path: 'Signup',
    component: SComponent
  },
  {
    path: 'ListingPage',
    component: ListingPageComponent
  },
  {
    path:'SellerPage',
    component: SellerPageComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
