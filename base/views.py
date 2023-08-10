from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.

def home(request):
    context = {'title' : 'Home'}
    return render(request, 'base/home.html', context)

def post(request):
    context = { 'title': 'Post'}
    return render(request, 'base/post.html', context)

def posts(request):
    context = {'title': 'Posts'}
    return render(request, 'base/posts.html', context)

def about_us(request):
    context = {'title': 'About us'}
    return render(request, 'base/about-us.html', context)

def profile(request):
    context = {'title': 'Profile'}
    return render(request, 'base/profile.html', context)