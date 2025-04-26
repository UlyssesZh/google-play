# Android 13

API 33

this seems to be the first version using `/fdfe/sync` only:

~~~
play -i com.android.vending -v 82941300
~~~

newest:

<http://dl.google.com/android/repository/sys-img/google_apis/x86_64-33_r17.zip>

oldest:

<http://dl.google.com/android/repository/sys-img/google_apis/x86_64-33_r05.zip>

MAKE SURE TO USE PIXEL 6

~~~
mitmproxy -s mitmproxy.py
emulator -avd Pixel_6_API_33 -writable-system

adb root
adb remount
adb reboot

adb root
adb remount

adb push com.android.vending-82941300.apk /system/priv-app
adb push privapp-permissions.xml /etc/permissions

adb push com.android.vending_45.9.19.apk /system/priv-app
adb push privapp-permissions.xml /etc/permissions

adb reboot
~~~

install system certificate

if you omit permission file, after reboot device will just load forever

https://source.android.com/docs/core/permissions/perms-allowlist
