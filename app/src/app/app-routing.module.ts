import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { LandingComponent } from 'src/app/page/landing/landing.component';
import { DatabaseComponent } from 'src/app/page/database/database.component';
import { RunningComponent } from 'src/app/page/running/running.component';
import { ScreenComponent } from 'src/app/page/screen/screen.component';

const routes: Routes = [
  { path: 'upload', component: DatabaseComponent },
  { path: 'running', component: RunningComponent },
  { path: 'screen', component: ScreenComponent },
  { path: '', redirectTo: '/running', pathMatch: 'full' },
  { path: '**', component: LandingComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
