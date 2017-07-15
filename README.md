// solidityHackathon

<link rel="stylesheet" type="text/css" href="semantic/dist/semantic.min.css">
<script src="jquery/dist/jquery.js">
	
</script>
<script src="semantic/dist/semantic.min.js"></script>

    //   -----------------------------------------------------------------------
    //   To GET active issues for a <USER>, <REPOSITORY>
    //       GET     https://api.github.com/repos/<USER>/<REPOSITORY>/issues?state=open
    //
    //   Save and track issue number
    //       <NUMBER> = response["number"]
    //   -----------------------------------------------------------------------
    //   Issue uniquely identified by <USER>, <REPOSITORY>, <NUMBER>

    //   -----------------------------------------------------------------------
    //   Poll issue to see if it is closed
    //       GET     https://api.github.com/repos/<USER>/<REPOSITORY>/issues/<NUMBER>
    //
    //       <CLOSED> = response["state"] == "closed"    (true/false)
    //       <TIMECLOSED> = response["closed_at"]
    //
    //   If closed, poll all issues and find one with the same timestamp
    //       GET     https://api.github.com/repos/<USER>/<REPOSITORY>/issues?state=closed
    //
    //       for <ISSUE> in response:
    //           <ISS-TIMECLOSED> = <ISSUE>["closed_at"]
    //           <ISS-NUMBER> = <ISSUE>["number"]
    //           if <ISS-TIMECLOSED> == <TIMECLOSED> and <ISS-NUMBER> != <NUMBER>:
    //               // This is the body containing the ethereum address to send ETH to
    //               <ETH-ADDR> = <ISSUE>["body"].parse(solidityHackathonApp:<ADDRESS>)
    //   -----------------------------------------------------------------------
