import { Component, OnInit } from '@angular/core';
import { SocketioService } from 'src/app/services/socketio.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-screen',
  templateUrl: './screen.component.html',
  styleUrls: ['./screen.component.scss']
})
export class ScreenComponent implements OnInit {

  date = new Date().getFullYear();
  pack: number = 1
  pack_remain: number = 0
  pack_total: number = 0
  receive_count: number = 0
  grad_total: number = 0
  received: any

  constructor(
    public socketService: SocketioService,
    public api: ApiService
  ) {
    this.get_all_grad()
  }

  ngOnInit(): void {
    this.socketService.getRunning().subscribe((ceremony: any) => {
      console.log(ceremony);
      if (this.grad_total == 0) {
        this.get_all_grad()
      } else {
        this.get_result(ceremony)
      }
    })
  }



  async get_all_grad() {
    await this.api.getAll("ceremonyall").then((res: any,rej?:any) => {
      if (res.status == 200) {
        this.grad_total = res.body.all_count
      }
    })
    await this.api.getBy("ceremony", this.pack).then((res: any,rej?:any) => {
      if (res.status == 200) {
        this.get_result(res.body)
      }
    })
  }

  async get_result(res: any) {
    this.pack_remain = res.pack_remain
    this.receive_count = res.receive_count
    this.pack_total = res.pack_count

    if (res.receive_result != undefined) {
      this.received = res.receive_result[0]
    } else {
      this.received = undefined
    }

    this.pack = res.ceremonypack
  }

}
