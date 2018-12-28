set-location D:\Projects\Rebel-L

$Shell = $Host.UI.RawUI
$size = $Shell.WindowSize
$size.width=120
$size.height=50
$Shell.WindowSize = $size
$size = $Shell.BufferSize
$size.width=120
$size.height=5000
$Shell.BufferSize = $size

$env:GOPATH="D:\Projects\Rebel-L\_goworkspace\"

function cdgop($project){
	cd $env:GOPATH\src\github.com\rebel-l\$project
}
