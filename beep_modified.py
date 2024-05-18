import sys
import frida

def main(target_process):

    pid = frida.spawn(target_process)
    session = frida.attach(pid)

  
    script = session.create_script("""
        console.log("|__> Starting Frida script");

        // Resolve the address of the Beep function in kernel32.dll
        var beepAddr = Module.findExportByName("kernel32.dll", "Beep");
        console.log("|__> Hooking Beep at " + beepAddr);

        Interceptor.attach(beepAddr, {
    onEnter: function(args) {
        console.log("|__> Called Beep");
        console.log("|__> Original Frequency: " + args[0].toInt32());
        console.log("|__> Original Duration: " + args[1].toInt32());

        // 修改频率和时间参数
        args[0] = ptr(2);  // 设置频率为2Hz
        args[1] = ptr(20);   // 设置时间为100ms

        console.log("|__> Modified Frequency: " + args[0].toInt32());
        console.log("|__> Modified Duration: " + args[1].toInt32());
    },
    onLeave: function(retval) {
        console.log("|__>|__> Returned from Beep");
    }
});
    """)


    script.load()


    frida.resume(pid)


    sys.stdin.read()


    session.detach()

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: %s <target_program>" % sys.argv[0])
        sys.exit(1)

    main(sys.argv[1])
