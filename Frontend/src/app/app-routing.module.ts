import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClassComponent } from './pages/class/class.component';
import { ResultsComponent } from './pages/results/results.component';
import { HomeComponent } from './pages/home/home.component';
import { AddClassComponent } from './pages/class/add-class/add-class.component';

const routes: Routes = [
  { path: 'results/:search', component: ResultsComponent },
  { path: 'class/:className', component: ClassComponent },
  { path: 'addClass', component: AddClassComponent },
  { path: '', component: HomeComponent, pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
  constructor() { }
}
