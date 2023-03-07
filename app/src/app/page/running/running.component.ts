import { Component, OnInit } from '@angular/core';
import { ApiService } from 'src/app/services/api.service';
import { SocketioService } from 'src/app/services/socketio.service';
import { Caremony } from 'src/app/models/caremony'

@Component({
  selector: 'app-running',
  templateUrl: './running.component.html',
  styleUrls: ['./running.component.scss'],
  providers: [SocketioService],
})
export class RunningComponent implements OnInit {

  pack: number = 1
  pack_remain: number = 0
  receive_count: number = 0
  pack_total: number = 0
  receive_result: Caremony[] = [];
  graduates_count: number = 0
  graduates_all = [] as any;
  remain_result: Caremony[] = [];

  socket_res: any;

  constructor(
    public api: ApiService,
    public socketService: SocketioService,
  ) {
    this.get_all_grad()
  }


  ngOnInit(): void {
    this.socket_res = this.socketService.getRunning().subscribe((ceremony: any) => {
      console.log(ceremony);
      this.get_result(ceremony)
    })
  }

  async get_all_grad() {
    await this.api.getAll("ceremonyall").then((res: any) => {
      this.graduates_count = res.all_count
      this.graduates_all = res.all_result

    })
    await this.api.getBy("ceremony", this.pack).then((res: any) => {
      this.get_result(res)
    })
  }

  async get_result(res: any) {
    this.remain_result = res.remain_result
    this.pack_remain = res.pack_remain
    this.receive_count = res.receive_count
    this.receive_result = res.receive_result
    this.pack_total = res.pack_count
  }

  async running_grad(studentcode: string, group: number) {
    this.pack = group
    this.api.put('ceremony', studentcode, "true").then((res: any) => {
      if (res.updated > 0) {
        this.socketService.sendRunning(this.pack.toString());
      }
    })
  }
  async return_grad(studentcode: string, group: number) {
    this.pack = group
    this.api.put('ceremony', studentcode, "false").then((res: any) => {
      if (res.updated > 0) {
        this.socketService.sendRunning(this.pack.toString());
      }
    })
  }
  async control_pack(action: string) {
    if (action == 'plus') {
      this.pack += 1
    } else {
      (this.pack) -= 1
    }
    this.refresh_data()
  }
  async refresh_data() {
    this.socketService.sendRunning(this.pack.toString());
  }
}
