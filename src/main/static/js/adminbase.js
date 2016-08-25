// JavaScript Document
function InitDataTable(Id,ajax_addr)
{
  return $(Id).dataTable({bFilter:true, ajax:ajax_addr,bPaginate:true,"bInfo": true,"serverSide": true,"lengthMenu": [[15, 10], [15,10]]});	
}

function IsNameNull(strName)
{
    curVal = $("[name="+strName+"]").val();
	if(curVal==null||curVal=="")
	{
		return true;
	}
	return false;
}
function IsIdNull(strId)
{
    curVal = $("#"+strId).val();
    if(curVal==null||curVal=="")
    {
        return true;
    }
    return false;
}
function WhileVal(fn,val,arr)
{
    	
	for(i=0;i<arr.length;i++)
	{
		if(fn(arr[i]) == val)
		{
			return arr[i];
		}
	}
	return null;
}