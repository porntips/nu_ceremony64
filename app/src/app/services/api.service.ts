import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { catchError, map, Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  base_path = 'http://localhost:8080/'

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": "access-control-allow-origin, access-control-allow-headers"

    })
  }
  constructor(
    public http: HttpClient,
  ) { }

  async post(path: string, body: any) {
    return new Promise((res, rej) => {
      this.http.post(this.base_path + path, JSON.stringify(body))
        .subscribe((data: any) => {
          res(data)
        }, (err: any) => {
          rej(err)
        });
    });
  }

  async put(path: string, studentcode: string, ceremony: string) {
    return new Promise((res, rej) => {
      this.http.put(`${this.base_path}${path}/${studentcode}/${ceremony}`, null)
        .subscribe((data: any) => {
          res(data)
        }, (err: any) => {
          rej(err)
        });
    });
  }
  async getAll(path: string) {
    return new Promise((res, rej) => {
      this.http.get(this.base_path + path)
        .subscribe((data: any) => {
          res(data)
        }, (err: any) => {
          rej(err)
        });
    });
  }
  async getBy(path: string, id:any) {
    return new Promise((res, rej) => {
      this.http.get(`${this.base_path}${path}/${id}`)
        .subscribe((data: any) => {
          res(data)
        }, (err: any) => {
          rej(err)
        });
    });
  }
}
