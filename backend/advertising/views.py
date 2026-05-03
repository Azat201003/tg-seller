from advertising.models import Advertise, Filter
from advertising.serializers import AdvertiseSerializer, FilterSerializer
from rest_framework import generics, permissions
from django.db.models import Count

class AdvertisesView(generics.ListCreateAPIView):
    queryset = Advertise.objects.prefetch_related("filters").all()
    serializer_class = AdvertiseSerializer
    permission_classes = [permissions.AllowAny]

    def get_queryset(self):
        filters = self.request.GET.getlist("filter")
        if len(filters) == 0:
          return super().get_queryset()
        
        advertisesQS = Filter.objects.filter(filter_id__in=filters).values('advertises').annotate(filtersCount=Count('advertises')).filter(filtersCount__gte=len(filters)//2).order_by("-filtersCount").values('advertises')
        print(filters)
        print("advertisesQS", advertisesQS.values_list("advertises"))

        advertises = Advertise.objects.filter(advertise_id__in=advertisesQS)

        return advertises

class FiltersView(generics.ListCreateAPIView):
    queryset = Filter.objects.all().filter(is_active=True).values("filter_id", "name")
    serializer_class = FilterSerializer
    permission_classes = [permissions.AllowAny]
