from django.contrib import admin
from django.urls import path
from advertising.views import AdvertisesView, FiltersView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('advertises/', AdvertisesView.as_view()),
    path('filters/', FiltersView.as_view())
]
