import { Component, OnInit } from '@angular/core';
import { ApiService } from 'src/app/services/api.service';
import { SocketioService } from 'src/app/services/socketio.service';
import { Caremony } from 'src/app/models/caremony'
import * as $ from 'jquery';

@Component({
  selector: 'app-running',
  templateUrl: './running.component.html',
  styleUrls: ['./running.component.scss'],
  providers: [SocketioService],
})
export class RunningComponent implements OnInit {

  pack: number = 1
  pack_max: number = 0
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

    $(document).on('keypress', function (e) {
      if (e.which == 13) {
        $("#runningBtn").click();
      } else if (e.which == 32) {
        $("#deleteBtn").click();
      }
    });
  }

  async get_all_grad() {
    await this.api.getAll("ceremonyall").then((res: any, rej?: any) => {
      if (res.status == 200) {
        this.graduates_count = res.body.all_count
        this.graduates_all = res.body.all_result
        if (res.body.all_result) {
          this.pack_max = Math.max.apply(Math, res.body.all_result.map(function (o: any) { return (Number(o.ceremonypack)); }));
        }
      }
    })
    await this.api.getBy("ceremony", this.pack).then((res: any, rej?: any) => {
      if (res.status == 200) {
        this.get_result(res.body)
      }
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
    this.api.put('ceremony', studentcode, "true").then((res: any, rej?: any) => {
      if (res.status == 200) {
        if (res.body.updated > 0) {
          this.socketService.sendRunning(this.pack.toString());
        }
      }
    })
  }
  async return_grad(studentcode: string, group: number) {
    this.pack = group
    this.api.put('ceremony', studentcode, "false").then((res: any, rej?: any) => {
      if (res.status == 200) {
        if (res.body.updated > 0) {
          this.socketService.sendRunning(this.pack.toString());
        }
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



  demo_test() {
    console.log(5678);

  }
}
