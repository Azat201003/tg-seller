from django.db import models

class Filter(models.Model):
    filter_id = models.AutoField(primary_key=True)
    name = models.CharField(blank=False, null=False, verbose_name="Название")
    is_active = models.BooleanField(blank=False, null=False, default=True, verbose_name="Активен")

    class Meta:
        verbose_name = "Фильтр"
        verbose_name_plural = "Фильтры"   

    def __str__(self):
        return f"{self.name} ({self.filter_id})"     

class Advertise(models.Model):
    advertise_id = models.AutoField(primary_key=True)
    link = models.CharField(blank=False, null=False, verbose_name="Ссылка")
    name = models.CharField(blank=False, null=False, verbose_name="Название")
    filters = models.ManyToManyField(Filter, related_name="advertises", verbose_name="Фильтры")
    
    class Meta:
        verbose_name = "Реклама"
        verbose_name_plural = "Рекламы"

    def __str__(self):
        return f"{self.name} ({self.advertise_id})"
