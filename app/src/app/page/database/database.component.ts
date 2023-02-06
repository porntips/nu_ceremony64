import { Component, OnInit, VERSION, ViewChild } from '@angular/core';
import { DatePipe } from '@angular/common';
import { ApiService } from 'src/app/services/api.service';
import { SocketioService } from 'src/app/services/socketio.service';

import Swal from 'sweetalert2';
import { Router } from '@angular/router';

export class CsvData {
  public studentcode: string | undefined;
  public sname: string | undefined;
  public degreecertificate: string | undefined;
  public facultyname: string | undefined;
  public hornor: number | undefined
  public ceremonygroup: number | undefined;
  public ceremonysequence: number | undefined;
  public ceremonydate: string | undefined;
  public ceremonypack: number | undefined;
  public ceremonypackno: number | undefined;
  public ceremonysex: string | undefined;
  public ceremonyprefix: string | undefined;
  public ceremony: boolean | undefined;
}

@Component({
  selector: 'app-database',
  templateUrl: './database.component.html',
  styleUrls: ['./database.component.scss']
})
export class DatabaseComponent implements OnInit {

  public records: any[] = [];
  @ViewChild('csvReader') csvReader: any;

  page: number = 1;
  count: number = 0;
  tableSize: number = 10;
  tableSizes: any = [10, 15, 20, 25];

  constructor(
    public datepipe: DatePipe,
    public api: ApiService,
    public socketService: SocketioService,
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  uploadListener($event: any): void {

    let text = [];
    let files = $event.srcElement.files;

    if (this.isValidCSVFile(files[0])) {

      let input = $event.target;
      let reader = new FileReader();
      reader.readAsText(input.files[0]);

      reader.onload = () => {
        let csvData = reader.result?.toString();
        let csvRecordsArray = (csvData!).split(/\r\n|\n/);

        let headersRow = this.getHeaderArray(csvRecordsArray);

        this.records = this.getDataRecordsArrayFromCSVFile(csvRecordsArray, headersRow.length);
      };

      reader.onerror = function () {
        console.log('error is occured while reading file!');
      };

    }
  }

  getDataRecordsArrayFromCSVFile(csvRecordsArray: any, headerLength: any) {
    let csvArr = [];

    for (let i = 1; i < csvRecordsArray.length; i++) {
      let curruntRecord = (csvRecordsArray[i]).split(',');
      if (curruntRecord.length == headerLength) {
        let csvRecord: CsvData = new CsvData();
        csvRecord.studentcode = curruntRecord[0].trim();
        csvRecord.sname = curruntRecord[1].trim();
        csvRecord.degreecertificate = curruntRecord[2].trim();
        csvRecord.facultyname = curruntRecord[3].trim();
        csvRecord.hornor = Number(curruntRecord[4].trim());
        csvRecord.ceremonygroup = Number(curruntRecord[5].trim());
        csvRecord.ceremonysequence = Number(curruntRecord[6].trim());
        csvRecord.ceremonydate = this.convertDate(curruntRecord[7].trim());
        csvRecord.ceremonypack = Number(curruntRecord[8].trim());
        csvRecord.ceremonypackno = Number(curruntRecord[9].trim());
        csvRecord.ceremonysex = curruntRecord[10].trim();
        csvRecord.ceremonyprefix = curruntRecord[11].trim();
        csvRecord.ceremony = Boolean(curruntRecord[12].trim());

        csvArr.push(csvRecord);
      }
    }
    return csvArr;
  }

  //check etension
  isValidCSVFile(file: any) {
    return file.name.endsWith(".csv");
  }

  getHeaderArray(csvRecordsArr: any) {
    let headers = (csvRecordsArr[0]).split(',');
    let headerArray = [];
    for (let j = 0; j < headers.length; j++) {
      headerArray.push(headers[j]);
    }
    return headerArray;
  }

  async uploadData() {
    await this.api.post("ceremony", this.records).then((res: any) => {
      this.socketService.sendRunning('1');

      Swal.fire({
        title: 'อัพโหลดสำเร็จ',
        text: `${(Number(res.created)).toLocaleString()} รายการ`,
        icon: 'success',
        confirmButtonText: `ตกลง`,
        showConfirmButton: false,
        timer: 2000
      }).then((e:any)=>{
        this.router.navigateByUrl('/running')
      })
    })
  }

  convertDate(str: string) {
    const [day, month, year] = str.split('/');
    const date = new Date(+year, +month - 1, +day);
    const d = this.datepipe.transform(date, 'yyyy-MM-dd')
    return d?.toString()
  }

  //pagination
  onTableDataChange(event: any) {
    this.page = event;
  }
  onTableSizeChange(event: any): void {
    this.tableSize = event.target.value;
    this.page = 1;
  }
}
