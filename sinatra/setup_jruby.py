
import subprocess
import sys
import setup_util

def start(args, logfile, errfile):
  setup_util.replace_text("sinatra/resin-web.xml", "mysql:\/\/.*:3306", "mysql://" + args.database_host + ":3306")

  try:
    subprocess.check_call("rvm jruby-1.7.4 do bundle install --gemfile=Gemfile-jruby", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    subprocess.check_call("cp Gemfile-jruby Gemfile", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    subprocess.check_call("cp Gemfile-jruby.lock Gemfile.lock", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    subprocess.check_call("rvm jruby-1.7.4 do warble war", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    subprocess.check_call("rm -rf $RESIN_HOME/webapps/*", shell=True, stderr=errfile, stdout=logfile)
    subprocess.check_call("cp sinatra.war $RESIN_HOME/webapps/sinatra.war", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    subprocess.check_call("$RESIN_HOME/bin/resinctl start", shell=True, stderr=errfile, stdout=logfile)
    return 0
  except subprocess.CalledProcessError:
    return 1
def stop(logfile, errfile):
  try:
    subprocess.check_call("$RESIN_HOME/bin/resinctl shutdown", shell=True, stderr=errfile, stdout=logfile)
    subprocess.check_call("rm Gemfile", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    subprocess.check_call("rm Gemfile.lock", shell=True, cwd="sinatra", stderr=errfile, stdout=logfile)
    return 0
  except subprocess.CalledProcessError:
    return 1
