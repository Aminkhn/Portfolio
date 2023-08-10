from django.urls import path
from . import views

urlpatterns = [
    path('', views.home, name = 'home'),
    path('post/', views.post, name = 'post'),
    path('posts/', views.posts, name = 'posts'),
    path('profile/', views.profile, name = 'profile'),
    path('about-us/', views.about_us, name = 'about-us')
]
