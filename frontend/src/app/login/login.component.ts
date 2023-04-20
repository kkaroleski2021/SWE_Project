import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'b-component',
  template: `
    <form class="box">
  <div class="field">
    <label class="label">Email</label>
    <div class="control">
      <input class="input" type="email" placeholder="e.g. alex@example.com">
    </div>
  </div>
  <div class="field">
    <label class="label">Password</label>
    <div class="control">
      <input class="input" type="password" placeholder="********">
    </div>
  </div>
  
  <button>Log In</button>

<nav>
  <ul>
  <li><button [routerLink]="['/Signup']">Don't have an account? Sign up now!</button></li>
  </ul>
  </nav>

</form>
  `,
  styles: [`
     `]
})

export class BComponent implements OnInit {

  constructor() {
    document.title = 'Login Page';
  }

  ngOnInit() {
  }

}

