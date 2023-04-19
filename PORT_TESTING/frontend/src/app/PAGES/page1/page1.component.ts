import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-page1',
  templateUrl: './page1.component.html',
  styleUrls: ['./page1.component.css']
})

export class Page1Component implements OnInit{
  public msg: any
  constructor( private http: HttpClient ) { }
  public ngOnInit(): void {
    this.getMessage()
  }

  getMessage(){
    let temp = JSON.parse(JSON.stringify(this.http.get('/page1')))
    console.log(temp)
  }
}


