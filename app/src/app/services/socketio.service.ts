import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { BehaviorSubject } from 'rxjs';
import * as io from "socket.io-client";

@Injectable({
  providedIn: 'root'
})
export class SocketioService {


  public socket: any;
  public grad$: BehaviorSubject<string> = new BehaviorSubject('');

  constructor() {
    // this.socket = io(environment.SOCKET_ENDPOINT, { forceNew: false })
    this.socket = io("http://localhost:8080")
  }

  public sendRunning(section: string) {
    this.socket.emit('ceremony', section);
  }
  public getRunning = () => {
    this.socket.on('graduate', async (grad: any) => {
      this.grad$.next(grad);
    });

    return this.grad$.asObservable();
  };

}
