import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { AdminComponent } from "./admin/admin.component";
import { HomeComponent } from "./home/home.component";
import { ProfileComponent } from "./profile/profile.component";
import { AuthComponent } from "./auth/auth.component";

export const routes: Routes = [
  { path: "", component: HomeComponent },
  { path: "login", component: AuthComponent },
  { path: "admin", component: AdminComponent },
  { path: "profile", component: ProfileComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
