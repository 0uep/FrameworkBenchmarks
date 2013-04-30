from django.db import models

# Create your models here.

class World(models.Model):
  randomnumber = models.IntegerField()
  class Meta:
    db_table = 'world'

class Fortune(models.Model):
  message = models.StringField()
  class Meta:
    db_table = 'fortune'
