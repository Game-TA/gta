using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class UIEventHandle : MonoBehaviour {
    public delegate void  DBottomButtonClick(MoveScrollViewContent.TABNUMBER tabNumber);
    public static event DBottomButtonClick EBottomButtonClick;
	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
		
	}
    //click events
    public void OnMessageButtonClick()
    {
        Debug.Log("bottom button message click");
        EBottomButtonClick(MoveScrollViewContent.TABNUMBER.ONE);
    }
    public void OnProjectButtonClick()
    {
        Debug.Log("bottom button project click");
        EBottomButtonClick(MoveScrollViewContent.TABNUMBER.TWO);
    }
    public void OnMacketButtonClick()
    {
        Debug.Log("bottom button macket click");
        EBottomButtonClick(MoveScrollViewContent.TABNUMBER.THREE);
    }
    public void OnMyButtonClick()
    {
        Debug.Log("bottom button my click");
        EBottomButtonClick(MoveScrollViewContent.TABNUMBER.FOUR);
    }
}
