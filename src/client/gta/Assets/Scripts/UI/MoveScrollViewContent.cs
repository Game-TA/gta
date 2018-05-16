using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class MoveScrollViewContent : MonoBehaviour {

    public enum TABNUMBER {ONE=1,TWO=2,THREE=3,FOUR=4 }
    public TABNUMBER tabNumer = TABNUMBER.ONE;
    public int tabWidth = 1080;

    public RectTransform content;
    private Vector2 contentInitSizeDelta;
    private void OnEnable()
    {
        UIEventHandle.EBottomButtonClick += MoveTab;
    }

    private void OnDisable()
    {
        UIEventHandle.EBottomButtonClick -= MoveTab;
    }

    // Use this for initialization
    void Start () {
        content = GetComponent<RectTransform>();
        if (checkRectTransfromComponent()) {
            contentInitSizeDelta = content.sizeDelta;
        }
    }
	
	// Update is called once per frame
	void Update () {
		
	}
    public void MoveTab(TABNUMBER _tabNumber) {
        if (checkRectTransfromComponent()) {
            Debug.Log("offsetMin="+content.offsetMin);
            Debug.Log("offsetMax="+content.offsetMax);
            Debug.Log("sizeDelta"+content.sizeDelta);
            content.offsetMin = new Vector2(tabWidth*(((int)_tabNumber)-1)*-1.0f, content.offsetMin.y);
            content.sizeDelta = contentInitSizeDelta;
        }
    }
    private bool checkRectTransfromComponent()
    {
        if (content == null)
        {
            Debug.Log("required RectTransform component on gameobject that name =" + gameObject.name);
            return false;
        }
        return true;
    }
}
