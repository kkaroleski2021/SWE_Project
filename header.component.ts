import { Component } from '@angular/core';

@Component({
  selector: 'app-header',
  template: `

    <div class="navbar is-white">
      <!-- logo -->
      <div class="navbar-brand">
        <a class ="navbar-item">
          <img src="assets/img/gator-logo.png">
      </a>
    </div>

    <div class="container content has-text-centered">
       <p> WELCOME TO SWAMPY SELLS </p>
    </div>
  </div>


  <nav class="breadcrumb is-centered" aria-label="breadcrumbs">
  <ul>
    <li><a href="#">Textbooks</a></li>
    <li><a href="#">Technology</a></li>
    <li><a href="#">Clothes</a></li>
    <li><a href="#">General Decor</a></li>
    <li><a href="#">Furniture</a></li>
    <li><a href="#">Tickets</a></li>
    <li class="is-active"><a href="#" aria-current="page">School Supplies</a></li>
  </ul>
</nav>

  `,
  styles: [`
    .container{
      font-family: "cool";
      font-size: 45px;
      align-items: center;
  }

  `]

})
export class HeaderComponent {

}
