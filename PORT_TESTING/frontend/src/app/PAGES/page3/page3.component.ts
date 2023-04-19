import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';


@Component({
  selector: 'app-page3',
  templateUrl: './page3.component.html',
  styleUrls: ['./page3.component.css']
})

export class Page3Component implements OnInit{
  public msg: any
  constructor( private http: HttpClient ) { }
  public ngOnInit(): void {
    this.msg = this.http.get("/page3").subscribe((response)=>{
      this.msg = response;
    })
  }
}

