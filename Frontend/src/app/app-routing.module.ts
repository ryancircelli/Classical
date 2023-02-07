import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClassComponent } from './pages/class/class.component';
import { ResultsComponent } from './pages/results/results.component';

const routes: Routes = [
  { path: 'results', component: ResultsComponent },
  { path: 'class', component: ClassComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
